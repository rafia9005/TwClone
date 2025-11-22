import { createRoot } from 'react-dom/client'
import "./styles/globals.css"
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Index from './app/index'

createRoot(document.getElementById('root')!).render(
  <BrowserRouter>
    <Routes>
      <Route path='' element={<Index />} />
    </Routes>
  </BrowserRouter>
)
