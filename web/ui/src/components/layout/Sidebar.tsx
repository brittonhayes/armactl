import { NavLink } from 'react-router-dom'
import { cn } from '@/lib/utils'
import { 
  Server, 
  Download, 
  Search, 
  Key, 
  Home,
  Settings
} from 'lucide-react'

const navigation = [
  { name: 'Dashboard', href: '/', icon: Home },
  { name: 'Server', href: '/server', icon: Server },
  { name: 'Mods', href: '/mods', icon: Download },
  { name: 'Steam Query', href: '/steam', icon: Search },
  { name: 'Keys', href: '/keys', icon: Key },
]

export function Sidebar() {
  return (
    <div className="flex h-screen w-64 flex-col bg-card border-r">
      {/* Logo/Header */}
      <div className="flex items-center justify-center h-16 border-b bg-primary text-primary-foreground">
        <div className="flex items-center space-x-2">
          <Settings className="h-6 w-6" />
          <span className="text-lg font-semibold">armactl</span>
        </div>
      </div>

      {/* Navigation */}
      <nav className="flex-1 p-4 space-y-1">
        {navigation.map((item) => (
          <NavLink
            key={item.name}
            to={item.href}
            className={({ isActive }) =>
              cn(
                'flex items-center space-x-3 px-3 py-2 rounded-md text-sm font-medium transition-colors',
                isActive
                  ? 'bg-primary text-primary-foreground'
                  : 'text-muted-foreground hover:bg-accent hover:text-accent-foreground'
              )
            }
          >
            <item.icon className="h-5 w-5 flex-shrink-0" />
            <span>{item.name}</span>
          </NavLink>
        ))}
      </nav>

      {/* Footer */}
      <div className="p-4 border-t text-xs text-muted-foreground">
        <p>ARMA 3 Server Manager</p>
      </div>
    </div>
  )
}