import { useAuth } from "@/context/auth"

export default function Layout() {
    const { user } = useAuth()
    console.log(user)
    return (
        <div>

        </div>
    )
}
