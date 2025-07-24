import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { useServerStatus } from '@/hooks/useServerStatus'

export function Header() {
  const { status, isLoading, startServer, stopServer } = useServerStatus()

  const handleServerToggle = async () => {
    if (status?.running) {
      await stopServer()
    } else {
      await startServer()
    }
  }

  return (
    <header className="h-16 border-b bg-card px-6 flex items-center justify-between">
      <div className="flex items-center space-x-4">
        <h1 className="text-xl font-semibold">ARMA Server Management</h1>
      </div>
      
      <div className="flex items-center space-x-4">
        {/* Server Status */}
        <div className="flex items-center space-x-2">
          <span className="text-sm text-muted-foreground">Status:</span>
          {isLoading ? (
            <Badge variant="secondary">Loading...</Badge>
          ) : (
            <Badge variant={status?.running ? "default" : "secondary"}>
              {status?.running ? "Running" : "Stopped"}
            </Badge>
          )}
        </div>
        
        {/* Quick Actions */}
        <div className="flex items-center space-x-2">
          <Button
            variant={status?.running ? "destructive" : "default"}
            size="sm"
            disabled={isLoading}
            onClick={handleServerToggle}
          >
            {status?.running ? "Stop" : "Start"} Server
          </Button>
        </div>
      </div>
    </header>
  )
}