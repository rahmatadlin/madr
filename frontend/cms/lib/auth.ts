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
          // Log more details for debugging
          const err = error as {
            message?: string;
            response?: { data?: unknown };
          };

          // Extract error message from response if available
          let errorMessage = err.message || "Login failed";

          // Try to extract error from axios response
          if (err.response?.data) {
            const responseData = err.response.data as {
              error?: string;
              message?: string;
            };
            errorMessage =
              responseData.error || responseData.message || errorMessage;
          }

          logger.error("Login error details", {
            message: errorMessage,
            originalError: err.message,
            response: err.response?.data,
            error: String(error),
          });

          // Throw error with specific message for NextAuth
          throw new Error(errorMessage);
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
        token.role = user.role;
      }
      return token;
    },
    async session({ session, token }) {
      if (session?.user && token) {
        session.user.accessToken = token.accessToken as string | undefined;
        session.user.role = token.role as string | undefined;
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
  debug: process.env.NODE_ENV === "development",
};
