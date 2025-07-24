import type { 
  ServerStatus, 
  ServerOptions, 
  ServerInfo,
  ModsListRequest,
  ModsListResponse,
  ModsSyncRequest,
  ValidationResult,
  KeysListResponse,
  ErrorResponse
} from '@/types/api'

class ApiClient {
  private baseUrl: string

  constructor(baseUrl: string = '/api') {
    this.baseUrl = baseUrl
  }

  private async request<T>(
    endpoint: string, 
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    }

    try {
      const response = await fetch(url, config)
      
      if (!response.ok) {
        const errorData: ErrorResponse = await response.json().catch(() => ({
          error: `HTTP ${response.status}`,
          code: response.status,
          message: response.statusText
        }))
        throw new Error(errorData.message || errorData.error)
      }

      return await response.json()
    } catch (error) {
      console.error(`API request failed: ${endpoint}`, error)
      throw error
    }
  }

  // Health check
  async health(): Promise<{ status: string; service: string }> {
    return this.request('/health')
  }

  // Server API
  async getServerStatus(): Promise<ServerStatus> {
    return this.request('/server/status')
  }

  async startServer(options: Partial<ServerOptions>): Promise<{ status: string; message: string }> {
    return this.request('/server/start', {
      method: 'POST',
      body: JSON.stringify(options)
    })
  }

  async stopServer(): Promise<{ status: string; message: string }> {
    return this.request('/server/stop', {
      method: 'POST'
    })
  }

  // Steam API
  async querySteamServer(address: string): Promise<ServerInfo> {
    return this.request(`/steam/query/${encodeURIComponent(address)}`)
  }

  // Mods API
  async listMods(request: ModsListRequest = {}): Promise<ModsListResponse> {
    const params = new URLSearchParams()
    if (request.directory) params.append('directory', request.directory)
    if (request.preset) params.append('preset', request.preset)
    
    const queryString = params.toString()
    return this.request(`/mods${queryString ? `?${queryString}` : ''}`)
  }

  async syncMods(request: ModsSyncRequest): Promise<{ status: string; message: string }> {
    return this.request('/mods/sync', {
      method: 'POST',
      body: JSON.stringify(request)
    })
  }

  async validateMods(directory: string, preset: string): Promise<ValidationResult> {
    return this.request('/mods/validate', {
      method: 'POST',
      body: JSON.stringify({ directory, preset })
    })
  }

  // Keys API
  async listKeys(directory?: string): Promise<KeysListResponse> {
    const params = new URLSearchParams()
    if (directory) params.append('directory', directory)
    
    const queryString = params.toString()
    return this.request(`/keys${queryString ? `?${queryString}` : ''}`)
  }

  async validateKeys(directory: string): Promise<ValidationResult> {
    return this.request('/keys/validate', {
      method: 'POST',
      body: JSON.stringify({ directory })
    })
  }
}

export const apiClient = new ApiClient()
export default apiClient