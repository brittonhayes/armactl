import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { useServerStatus } from '@/hooks/useServerStatus'
import { Server, Users, HardDrive, Clock } from 'lucide-react'

export function Dashboard() {
  const { status, isLoading, startServer, stopServer } = useServerStatus()

  const handleServerToggle = async () => {
    if (status?.running) {
      await stopServer()
    } else {
      await startServer()
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
        <p className="text-muted-foreground">
          Overview of your ARMA 3 server status and quick actions
        </p>
      </div>

      {/* Status Cards */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Server Status</CardTitle>
            <Server className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              <Badge variant={status?.running ? "default" : "secondary"}>
                {isLoading ? "Loading..." : (status?.running ? "Running" : "Stopped")}
              </Badge>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Players</CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{status?.players || "0"}</div>
            <p className="text-xs text-muted-foreground">
              / 64 max
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Uptime</CardTitle>
            <Clock className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{status?.uptime || "0m"}</div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Memory Usage</CardTitle>
            <HardDrive className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{status?.memory || "0"} MB</div>
          </CardContent>
        </Card>
      </div>

      {/* Quick Actions */}
      <Card>
        <CardHeader>
          <CardTitle>Quick Actions</CardTitle>
          <CardDescription>
            Common server management tasks
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex space-x-4">
            <Button 
              variant={status?.running ? "destructive" : "default"}
              disabled={isLoading}
              onClick={handleServerToggle}
            >
              {status?.running ? "Stop Server" : "Start Server"}
            </Button>
            <Button variant="outline" disabled={!status?.running}>
              Restart Server
            </Button>
            <Button variant="outline">
              View Logs
            </Button>
          </div>
        </CardContent>
      </Card>

      {/* Recent Activity */}
      <Card>
        <CardHeader>
          <CardTitle>Recent Activity</CardTitle>
          <CardDescription>
            Latest server events and changes
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-2 text-sm">
            <div className="flex justify-between">
              <span>Server started</span>
              <span className="text-muted-foreground">2 hours ago</span>
            </div>
            <div className="flex justify-between">
              <span>Mods synchronized</span>
              <span className="text-muted-foreground">3 hours ago</span>
            </div>
            <div className="flex justify-between">
              <span>Configuration updated</span>
              <span className="text-muted-foreground">1 day ago</span>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}