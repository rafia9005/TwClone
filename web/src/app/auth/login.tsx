import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Form, FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useState } from "react"
import { metaData } from "@/content"
import Cookies from "js-cookie"
import { authAPI } from "@/lib/api"
import { Feather } from "lucide-react"

const loginSchema = z.object({
  identifier: z.string().min(1, "Email or username is required"),
  password: z.string().min(1, "Password is required"),
})

type LoginValues = z.infer<typeof loginSchema>

export default function Login() {
  const [apiError, setApiError] = useState<string | null>(null)
  const form = useForm<LoginValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: { identifier: "", password: "" },
    mode: "onTouched",
  })

  async function onSubmit(values: LoginValues) {
    setApiError(null)
    try {
      const loginData = values.identifier.includes("@")
        ? { email: values.identifier, password: values.password }
        : { username: values.identifier, password: values.password }
      
      const response = await authAPI.login(loginData)
      
      // Save token
      if (response.token) {
        Cookies.set("accessToken", response.token, { path: "/", expires: 7 })
        location.href = "/"
      }
    } catch (err: any) {
      setApiError(err?.response?.data?.message || "Login failed")
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-background px-4">
      <div className="w-full max-w-md">
        <Card>
          <CardHeader className="items-center">
            <div className="mb-2 flex items-center justify-center">
              <svg viewBox="0 0 32 32" width={40} height={40} className="text-blue-500" fill="currentColor"><g><path d="M31.94 6.1a1.13 1.13 0 0 0-1.18-.28l-4.13 1.42a.47.47 0 0 1-.56-.19 11.13 11.13 0 0 0-2.1-2.44A10.93 10.93 0 0 0 16 2.67a10.93 10.93 0 0 0-8.01 3.94 11.13 11.13 0 0 0-2.1 2.44.47.47 0 0 1-.56.19L1.2 5.82A1.13 1.13 0 0 0 0 7.13v17.74a1.13 1.13 0 0 0 1.2 1.31l4.13-1.42a.47.47 0 0 1 .56.19 11.13 11.13 0 0 0 2.1 2.44A10.93 10.93 0 0 0 16 29.33a10.93 10.93 0 0 0 8.01-3.94 11.13 11.13 0 0 0 2.1-2.44.47.47 0 0 1 .56-.19l4.13 1.42A1.13 1.13 0 0 0 32 24.87V7.13a1.13 1.13 0 0 0-.06-.33ZM16 27.33A9.33 9.33 0 1 1 25.33 18 9.34 9.34 0 0 1 16 27.33Z"/></g></svg>
            </div>
            <CardTitle className="text-2xl text-center">Sign in to {metaData.title}</CardTitle>
            <CardDescription className="text-center">Welcome back — sign in to continue</CardDescription>
          </CardHeader>
          <CardContent>
            <Form {...form}>
              <form className="space-y-4" onSubmit={form.handleSubmit(onSubmit)}>
                {apiError && (
                  <div className="mb-2 text-sm text-red-600 text-center">{apiError}</div>
                )}
                <FormField
                  control={form.control}
                  name="identifier"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Email or username</FormLabel>
                      <FormControl>
                        <Input placeholder="you@example.com or username" autoComplete="username" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl>
                        <Input type="password" placeholder="••••••••" autoComplete="current-password" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <div className="flex items-center justify-between">
                  <label className="flex items-center gap-2 text-sm text-muted-foreground">
                    <input type="checkbox" className="rounded" /> Remember me
                  </label>
                  <a href="#" className="text-sm text-blue-600 hover:underline">Forgot?</a>
                </div>
                <Button className="w-full mt-2" type="submit">Sign in</Button>
              </form>
            </Form>
            <p className="mt-6 text-center text-sm text-muted-foreground">
              New to TwClone?{' '}
              <a href="/register" className="text-blue-600 hover:underline font-medium">Create an account</a>
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
