import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

export function ServerPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Server Management</h1>
        <p className="text-muted-foreground">
          Configure and control your ARMA 3 dedicated server
        </p>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Server Configuration</CardTitle>
          <CardDescription>
            Configure server settings and startup options
          </CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">Server configuration will be implemented here.</p>
        </CardContent>
      </Card>
    </div>
  )
}