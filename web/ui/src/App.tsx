import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Layout } from '@/components/layout/Layout'
import { Dashboard } from '@/pages/Dashboard'
import { ServerPage } from '@/pages/ServerPage'
import { ModsPage } from '@/pages/ModsPage'
import { SteamPage } from '@/pages/SteamPage'
import { KeysPage } from '@/pages/KeysPage'

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/server" element={<ServerPage />} />
          <Route path="/mods" element={<ModsPage />} />
          <Route path="/steam" element={<SteamPage />} />
          <Route path="/keys" element={<KeysPage />} />
        </Routes>
      </Layout>
    </Router>
  )
}

export default App