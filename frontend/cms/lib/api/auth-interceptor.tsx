"use client";

import { useEffect } from "react";
import { useSession } from "next-auth/react";
import { setAuthToken } from "./client";

export function AuthInterceptor({ children }: { children: React.ReactNode }) {
  const { data: session } = useSession();

  useEffect(() => {
    const token = (session?.user as any)?.accessToken;
    setAuthToken(token || null);
  }, [session]);

  return <>{children}</>;
}

