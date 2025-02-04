import { IconType } from "react-icons/lib";

import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import React from "react";

type ProviderLoginBtnProps = {
    LogoIcon: IconType;
    providerName: string;
};

const ProviderLoginBtn = React.forwardRef<
    HTMLAnchorElement,
    React.ComponentProps<"a"> & ProviderLoginBtnProps
>(({ className, href, LogoIcon, providerName, ...props }, ref) => {
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
            ref={ref}
            {...props}
        >
            <span className="flex items-center gap-2">{providerName}</span>
            <LogoIcon
                style={{
                    width: "1.5em",
                    height: "1.5em",
                }}
            />
        </a>
    );
});

ProviderLoginBtn.displayName = "Anchor";

export default ProviderLoginBtn;
