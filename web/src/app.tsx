import { createRoot } from 'react-dom/client'
import "./styles/globals.css"
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Index from './app/index'
import Register from './app/auth/register'
import Login from './app/auth/login'
import { ThemeProvider } from './components/theme-provider'

createRoot(document.getElementById('root')!).render(
  <ThemeProvider defaultTheme='system'>
    <BrowserRouter>
      <Routes>
        <Route path='' element={<Index />} />
        <Route path='register' element={<Register />} />
        <Route path='login' element={<Login />} />
      </Routes>
    </BrowserRouter>
  </ThemeProvider>
)
