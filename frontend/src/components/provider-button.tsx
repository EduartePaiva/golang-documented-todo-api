import { IconType } from "react-icons/lib";

import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";

type ProviderLoginBtnProps = {
    LogoIcon: IconType;
    name: string;
    href: string;
    className?: string;
};

export default function ProviderLoginBtn({
    LogoIcon,
    name,
    href,
    className,
}: ProviderLoginBtnProps) {
    return (
        <a
            className={cn(
                buttonVariants({
                    size: "sm",
                    variant: "ghost",
                }),
                className
            )}
            type="submit"
            href={href}
        >
            <span className="flex items-center gap-2">{name}</span>
            <LogoIcon
                style={{
                    width: "1.5em",
                    height: "1.5em",
                }}
            />
        </a>
    );
}
