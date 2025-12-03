/**
 * Centralized logging service
 * Provides consistent logging across the application
 */

type LogLevel = "debug" | "info" | "warn" | "error";

class Logger {
  private isDevelopment = process.env.NODE_ENV === "development";

  private log(level: LogLevel, message: string, ...args: unknown[]) {
    if (!this.isDevelopment && level === "debug") {
      return;
    }

    const timestamp = new Date().toISOString();
    const prefix = `[${timestamp}] [${level.toUpperCase()}]`;

    switch (level) {
      case "debug":
        console.debug(prefix, message, ...args);
        break;
      case "info":
        console.info(prefix, message, ...args);
        break;
      case "warn":
        console.warn(prefix, message, ...args);
        break;
      case "error":
        console.error(prefix, message, ...args);
        break;
    }
  }

  debug(message: string, ...args: unknown[]) {
    this.log("debug", message, ...args);
  }

  info(message: string, ...args: unknown[]) {
    this.log("info", message, ...args);
  }

  warn(message: string, ...args: unknown[]) {
    this.log("warn", message, ...args);
  }

  error(message: string, ...args: unknown[]) {
    this.log("error", message, ...args);
  }

  api(method: string, url: string, data?: unknown) {
    this.debug(`API ${method} ${url}`, data);
  }
}

export const logger = new Logger();
