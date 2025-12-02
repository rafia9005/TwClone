import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Form, FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { metaData } from "@/content"
import { authAPI } from "@/lib/api"

const registerSchema = z.object({
    email: z.string().min(1, "Email is required").email("Invalid email address"),
    name: z.string().min(1, "Name is required"),
    username: z.string().min(1, "Username is required"),
    password: z.string().min(6, "Password must be at least 6 characters"),
})

type registerValues = z.infer<typeof registerSchema>

export default function Register() {
    const navigate = useNavigate()
    const [apiError, setApiError] = useState<string | null>(null)
    const form = useForm<registerValues>({
        resolver: zodResolver(registerSchema),
        defaultValues: { email: "", name: "", username: "", password: "" },
        mode: "onTouched",
    })

    async function onSubmit(values: registerValues) {
        setApiError(null)
        try {
            await authAPI.register(values)
            navigate("/login")
        } catch (err: any) {
            const apiErr = err?.response?.data
            if (apiErr?.errors && Array.isArray(apiErr.errors)) {
                setApiError(apiErr.errors.map((e: any) => e.message).join(", "))
            } else {
                setApiError(apiErr?.message || "Registration failed")
            }
        }
    }

    return (
        <div className="min-h-screen flex items-center justify-center bg-background px-4">
            <div className="w-full max-w-md">
                <Card>
                    <CardHeader className="items-center">
                        <div className="mb-2 flex items-center justify-center">
                            <svg viewBox="0 0 32 32" width={40} height={40} className="text-blue-500" fill="currentColor"><g><path d="M31.94 6.1a1.13 1.13 0 0 0-1.18-.28l-4.13 1.42a.47.47 0 0 1-.56-.19 11.13 11.13 0 0 0-2.1-2.44A10.93 10.93 0 0 0 16 2.67a10.93 10.93 0 0 0-8.01 3.94 11.13 11.13 0 0 0-2.1 2.44.47.47 0 0 1-.56.19L1.2 5.82A1.13 1.13 0 0 0 0 7.13v17.74a1.13 1.13 0 0 0 1.2 1.31l4.13-1.42a.47.47 0 0 1 .56.19 11.13 11.13 0 0 0 2.1 2.44A10.93 10.93 0 0 0 16 29.33a10.93 10.93 0 0 0 8.01-3.94 11.13 11.13 0 0 0 2.1-2.44.47.47 0 0 1 .56-.19l4.13 1.42A1.13 1.13 0 0 0 32 24.87V7.13a1.13 1.13 0 0 0-.06-.33ZM16 27.33A9.33 9.33 0 1 1 25.33 18 9.34 9.34 0 0 1 16 27.33Z" /></g></svg>
                        </div>
                        <CardTitle className="text-2xl text-center">Sign up to {metaData.title}</CardTitle>
                        <CardDescription className="text-center">Welcome back — sign up to continue</CardDescription>
                    </CardHeader>
                    <CardContent>
                        <Form {...form}>
                            <form className="space-y-4" onSubmit={form.handleSubmit(onSubmit)}>
                                {apiError && (
                                    <div className="mb-2 text-sm text-red-600 text-center">{apiError}</div>
                                )}
                                <FormField
                                    control={form.control}
                                    name="email"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Email</FormLabel>
                                            <FormControl>
                                                <Input placeholder="you@example.com or username" autoComplete="username" {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />

                                <FormField
                                    control={form.control}
                                    name="name"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Name</FormLabel>
                                            <FormControl>
                                                <Input type="text" placeholder="Your name" autoComplete="name" {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />

                                <FormField
                                    control={form.control}
                                    name="username"
                                    render={({ field }) => (
                                        <FormItem>
                                            <FormLabel>Username</FormLabel>
                                            <FormControl>
                                                <Input type="text" placeholder="Your username" autoComplete="username" {...field} />
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
                                                <Input type="password" placeholder="Your password" autoComplete="password" {...field} />
                                            </FormControl>
                                            <FormMessage />
                                        </FormItem>
                                    )}
                                />
                                <Button className="w-full mt-2" type="submit">Sign up</Button>
                            </form>
                        </Form>
                        <p className="mt-6 text-center text-sm text-muted-foreground">
                            Veteran in {metaData.title}?{' '}
                            <a href="/login" className="text-blue-600 hover:underline font-medium">Login to your account</a>
                        </p>
                    </CardContent>
                </Card>
            </div>
        </div>
    )
}
