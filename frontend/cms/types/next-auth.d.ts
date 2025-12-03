import "next-auth";
import "next-auth/jwt";

declare module "next-auth" {
  interface User {
    role?: string;
    accessToken?: string;
    refreshToken?: string;
  }

  interface Session {
    user: {
      id?: string;
      name?: string | null;
      email?: string | null;
      role?: string;
      accessToken?: string;
    };
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    accessToken?: string;
    refreshToken?: string;
    role?: string;
  }
}

