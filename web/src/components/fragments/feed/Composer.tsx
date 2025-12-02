import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { useState, useRef } from "react"
import { mediaAPI } from "@/lib/api"
import { Image, X, Smile } from "lucide-react"

export default function Composer({ user, onTweet }: { user?: any; onTweet: (text: string, mediaUrl?: string) => void }) {
  const [text, setText] = useState("")
  const [mediaPreview, setMediaPreview] = useState<string | null>(null)
  const [mediaFile, setMediaFile] = useState<File | null>(null)
  const [uploading, setUploading] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)
  
  const disabled = !user
  const max = 280

  const handleImageSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      if (file.size > 5 * 1024 * 1024) {
        alert("Image must be less than 5MB")
        return
      }
      
      setMediaFile(file)
      const reader = new FileReader()
      reader.onloadend = () => {
        setMediaPreview(reader.result as string)
      }
      reader.readAsDataURL(file)
    }
  }

  const removeImage = () => {
    setMediaPreview(null)
    setMediaFile(null)
    if (fileInputRef.current) {
      fileInputRef.current.value = ""
    }
  }

  const handleTweet = async () => {
    if (text.trim().length === 0 || uploading) return
    
    let mediaUrl: string | undefined = undefined
    
    try {
      // Upload media if present
      if (mediaFile) {
        setUploading(true)
        const media = await mediaAPI.upload(mediaFile)
        mediaUrl = media.media_url
      }
      
      // Post tweet with media URL
      onTweet(text.trim(), mediaUrl)
      
      // Reset state
      setText("")
      removeImage()
    } catch (error) {
      console.error("Failed to post tweet:", error)
      alert("Failed to post tweet. Please try again.")
    } finally {
      setUploading(false)
    }
  }

  return (
    <div className="border-b p-4">
      <div className="flex gap-3">
        <Avatar className="h-12 w-12">
          {user?.avatar ? (
            <AvatarImage src={user.avatar} alt={user.name} />
          ) : (
            <AvatarFallback>{user?.name?.[0] ?? "G"}</AvatarFallback>
          )}
        </Avatar>
        
        <div className="flex-1">
          <textarea
            className="w-full min-h-[80px] resize-none border-0 focus:outline-none text-base placeholder:text-muted-foreground bg-transparent"
            placeholder={disabled ? "Sign in to post" : "What's happening?"}
            value={text}
            onChange={(e) => setText(e.target.value)}
            disabled={disabled || uploading}
          />

          {/* Image Preview */}
          {mediaPreview && (
            <div className="relative mt-3 rounded-2xl overflow-hidden">
              <img 
                src={mediaPreview} 
                alt="Upload preview" 
                className="w-full max-h-96 object-cover"
              />
              <button
                onClick={removeImage}
                className="absolute top-2 right-2 p-2 bg-black/60 hover:bg-black/80 rounded-full transition-colors"
                disabled={uploading}
              >
                <X className="w-4 h-4 text-white" />
              </button>
            </div>
          )}

          <div className="mt-3 flex items-center justify-between border-t pt-3">
            <div className="flex items-center gap-2">
              <input
                ref={fileInputRef}
                type="file"
                accept="image/*"
                onChange={handleImageSelect}
                className="hidden"
                disabled={disabled || uploading || !!mediaPreview}
              />
              
              <Button
                variant="ghost"
                size="icon"
                className="h-9 w-9 text-primary hover:bg-primary/10"
                onClick={() => fileInputRef.current?.click()}
                disabled={disabled || uploading || !!mediaPreview}
                title="Add image"
              >
                <Image className="w-5 h-5" />
              </Button>
              
              <Button
                variant="ghost"
                size="icon"
                className="h-9 w-9 text-primary hover:bg-primary/10"
                disabled={disabled || uploading}
                title="Add emoji"
              >
                <Smile className="w-5 h-5" />
              </Button>
            </div>

            <div className="flex items-center gap-3">
              <div className={`text-sm ${text.length > max ? "text-red-500" : "text-muted-foreground"}`}>
                {text.length > 0 && `${text.length}/${max}`}
              </div>
              
              <Button
                disabled={text.trim().length === 0 || text.length > max || disabled || uploading}
                onClick={handleTweet}
                size="sm"
                className="rounded-full font-bold px-4"
              >
                {uploading ? "Posting..." : "Post"}
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
