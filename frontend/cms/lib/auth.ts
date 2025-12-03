import type { NextAuthConfig } from "next-auth";
import Credentials from "next-auth/providers/credentials";
import { authApi } from "./api/auth";
import { logger } from "./logger";

export const authOptions: NextAuthConfig = {
  providers: [
    Credentials({
      name: "Credentials",
      credentials: {
        username: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        if (!credentials?.username || !credentials?.password) {
          logger.warn("Login attempt with missing credentials");
          return null;
        }

        try {
          const username = String(credentials.username);
          const password = String(credentials.password);
          
          const response = await authApi.login({
            username,
            password,
          });

          logger.info(`User ${response.user.username} logged in successfully`);

          return {
            id: response.user.id.toString(),
            name: response.user.username,
            email: response.user.email,
            role: response.user.role,
            accessToken: response.access_token,
            refreshToken: response.refresh_token,
          };
        } catch (error) {
          logger.error("Login failed", error);
          return null;
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.accessToken = (user as any).accessToken;
        token.refreshToken = (user as any).refreshToken;
        token.role = (user as any).role;
      }
      return token;
    },
    async session({ session, token }) {
      if (session.user) {
        (session.user as any).accessToken = token.accessToken;
        (session.user as any).role = token.role;
      }
      return session;
    },
  },
  pages: {
    signIn: "/login",
  },
  session: {
    strategy: "jwt",
  },
  secret: process.env.NEXTAUTH_SECRET || "your-secret-key-change-in-production",
};

