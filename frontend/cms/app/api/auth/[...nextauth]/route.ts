import NextAuth from "next-auth";
import { authOptions } from "@/lib/auth";

const handler = NextAuth(authOptions);

// NextAuth v5 beta compatibility with Next.js 16
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const GET = handler as any;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const POST = handler as any;

