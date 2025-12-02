"use client"

import { Search, MoreHorizontal } from "lucide-react"

export default function RightSidebar() {
  return (
    <aside className="w-80 p-4 hidden lg:block sticky top-0 h-screen">
      <div className="relative mb-4">
        <Search className="absolute left-4 top-3 text-muted-foreground" size={20} />
        <input
          type="text"
          placeholder="Search"
          className="w-full bg-secondary rounded-full pl-12 pr-4 py-3 outline-none text-foreground placeholder-muted-foreground"
        />
      </div>

      <div className="bg-secondary rounded-2xl p-4 mb-4">
        <h3 className="text-xl font-bold mb-4">What's happening!?</h3>

        <div className="space-y-4">
          {[
            { category: "Technology • Trending", title: "TypeScript", posts: "145K posts" },
            { category: "Technology • Live", title: "React Conf", posts: "89K posts" },
            { category: "Design • Popular", title: "Web Design", posts: "56K posts" },
            { category: "Development • Trending", title: "#WebDevelopment", posts: "234K posts" },
          ].map((item, i) => (
            <div
              key={i}
              className="hover:bg-secondary/50 p-3 rounded-xl cursor-pointer transition-colors flex justify-between items-start"
            >
              <div>
                <div className="text-sm text-muted-foreground">{item.category}</div>
                <div className="font-bold text-base">{item.title}</div>
                <div className="text-sm text-muted-foreground">{item.posts}</div>
              </div>
              <MoreHorizontal size={18} className="text-muted-foreground" />
            </div>
          ))}
        </div>
      </div>

      <div className="bg-secondary rounded-2xl p-4">
        <h3 className="text-xl font-bold mb-4">Subscribe to Pro</h3>
        <p className="text-muted-foreground mb-4">
          Subscribe to unlock new features and if eligible, receive a share of ads revenue.
        </p>
        <button className="w-full bg-primary hover:bg-primary/90 text-primary-foreground font-bold py-2 rounded-full transition-all">
          Subscribe
        </button>
      </div>
    </aside>
  )
}
