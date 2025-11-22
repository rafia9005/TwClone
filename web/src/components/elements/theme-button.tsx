import { Moon, Sun } from "lucide-react"
import { useTheme } from "@/components/theme-provider"
import { Button } from "@/components/ui/button"

type ThemeToggleProps = {
  variant?: "default" | "rounded"
}

export function ThemeButton({ variant = "default" }: ThemeToggleProps) {
  const { theme, setTheme } = useTheme()

  const toggleTheme = () => {
    setTheme(theme === "dark" ? "light" : "dark")
  }

  return (
    <Button
      variant={variant === "rounded" ? "ghost" : "outline"}
      size={variant === "rounded" ? "icon" : "default"}
      aria-label={`Switch to ${theme === "dark" ? "light" : "dark"} mode`}
      onClick={toggleTheme}
      className={variant === "rounded" ? "rounded-full border p-2" : "rounded-lg"}
    >
      {theme === "dark" ? (
        <Sun className="h-5 w-5 transition-transform duration-200" />
      ) : (
        <Moon className="h-5 w-5 transition-transform duration-200" />
      )}
    </Button>
  )
}