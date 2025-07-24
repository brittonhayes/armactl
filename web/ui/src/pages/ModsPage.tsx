import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

export function ModsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Mod Management</h1>
        <p className="text-muted-foreground">
          Manage ARMA 3 modifications and presets
        </p>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Mod List</CardTitle>
          <CardDescription>
            View and manage installed modifications
          </CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">Mod management will be implemented here.</p>
        </CardContent>
      </Card>
    </div>
  )
}