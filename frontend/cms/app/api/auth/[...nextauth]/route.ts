import NextAuth from "next-auth";
import { authOptions } from "@/lib/auth";

// NextAuth v5 beta - NextAuth returns handlers object with GET and POST methods
const { handlers } = NextAuth(authOptions);

// Export GET and POST handlers
export const GET = handlers.GET;
export const POST = handlers.POST;
