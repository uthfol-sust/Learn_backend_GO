import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './styles/gobal.css'
import App from './App.jsx'
import { AuthProvider } from './context/AuthContext.jsx'
import { SidebarProvider } from './context/SidebarContext.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <AuthProvider>
      <SidebarProvider>
        <App />
      </SidebarProvider>
    </AuthProvider>
  </StrictMode>,
)
