import { useState, useEffect } from 'react'
import type { ServerStatus } from '@/types/api'
import { apiClient } from '@/lib/api'

export function useServerStatus() {
  const [status, setStatus] = useState<ServerStatus | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchStatus = async () => {
      try {
        setIsLoading(true)
        const data = await apiClient.getServerStatus()
        setStatus(data)
        setError(null)
      } catch (err) {
        setError(err instanceof Error ? err.message : 'An error occurred')
        // Set default status when API is not available
        setStatus({
          running: false,
          players: 0,
          memory: 0,
          uptime: '0m'
        })
      } finally {
        setIsLoading(false)
      }
    }

    fetchStatus()
    
    // Poll for status updates every 30 seconds
    const interval = setInterval(fetchStatus, 30000)
    
    return () => clearInterval(interval)
  }, [])

  const startServer = async (options: any = {}) => {
    try {
      await apiClient.startServer(options)
      // Refresh status after starting
      const data = await apiClient.getServerStatus()
      setStatus(data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to start server')
    }
  }

  const stopServer = async () => {
    try {
      await apiClient.stopServer()
      // Refresh status after stopping
      const data = await apiClient.getServerStatus()
      setStatus(data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to stop server')
    }
  }

  return { 
    status, 
    isLoading, 
    error,
    startServer,
    stopServer
  }
}