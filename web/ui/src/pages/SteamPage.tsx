import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

export function SteamPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Steam Server Query</h1>
        <p className="text-muted-foreground">
          Query Steam servers for information and player counts
        </p>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Server Query</CardTitle>
          <CardDescription>
            Query Steam servers for real-time information
          </CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">Steam server query will be implemented here.</p>
        </CardContent>
      </Card>
    </div>
  )
}