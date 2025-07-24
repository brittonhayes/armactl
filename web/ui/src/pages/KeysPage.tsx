import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

export function KeysPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Key Management</h1>
        <p className="text-muted-foreground">
          Manage bikey files for ARMA modifications
        </p>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Key Files</CardTitle>
          <CardDescription>
            View and manage .bikey files for mod verification
          </CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">Key management will be implemented here.</p>
        </CardContent>
      </Card>
    </div>
  )
}