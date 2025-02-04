import { LogOut } from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "./ui/dropdown-menu";

interface UserAvatarProps {
    avatar_url: string;
    name: string;
}

export default function UserAvatar({ avatar_url, name }: UserAvatarProps) {
    return (
        <div>
            <DropdownMenu dir="ltr">
                <DropdownMenuTrigger asChild>
                    <Avatar className="h-9 w-9 cursor-pointer transition hover:scale-105">
                        <AvatarImage src={avatar_url} alt="avatar" />
                        <AvatarFallback>
                            <span className="font-mono font-semibold">
                                {name.substring(0, 2)}
                            </span>
                        </AvatarFallback>
                    </Avatar>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                    <DropdownMenuLabel>User Options</DropdownMenuLabel>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem asChild>
                        <a
                            className="flex cursor-pointer items-center gap-2"
                            href="/api/v1/login/logout"
                        >
                            <LogOut />
                            Log Out
                        </a>
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    );
}
