import { createRoot } from 'react-dom/client'
import "./styles/globals.css"
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Index from './app/index'
import Register from './app/auth/register'
import Login from './app/auth/login'
import ProfilePage from './app/profile/[id]'
import NotificationsPage from './app/notifications/index'
import ExplorePage from './app/explore/index'
import TweetDetailPage from './app/tweet/[id]'
import { ThemeProvider } from './components/theme-provider'
import { AuthProvider } from './context/auth'

// Optional: provide a developer test user via Vite env variable VITE_DEV_USER_JSON
// Example (in .env.development):
// VITE_DEV_USER_JSON='{"id":1,"email":"rafia9005@gmail.com","name":"Ahmad Rafi\'i","username":"rafia9005","created_at":"2025-11-26T18:24:49+07:00","updated_at":"2025-11-26T18:24:49+07:00"}'
const devUser = import.meta.env.VITE_DEV_USER_JSON ? JSON.parse(import.meta.env.VITE_DEV_USER_JSON as string) : undefined

createRoot(document.getElementById('root')!).render(
  <AuthProvider initialUser={devUser}>
    <ThemeProvider defaultTheme='system'>
      <BrowserRouter>
        <Routes>
          <Route path='' element={<Index />} />
          <Route path='register' element={<Register />} />
          <Route path='login' element={<Login />} />
          <Route path='profile/:id' element={<ProfilePage />} />
          <Route path='notifications' element={<NotificationsPage />} />
          <Route path='explore' element={<ExplorePage />} />
          <Route path='tweet/:id' element={<TweetDetailPage />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  </AuthProvider>
)
