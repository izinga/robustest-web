import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/layout/Header";
import Footer from "@/components/layout/Footer";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
});

export const metadata: Metadata = {
  title: "RobusTest - Enterprise Mobile Device Lab Management",
  description: "Your Devices. Your Data. Your Control. Enterprise-grade mobile device lab management deployed on your premises. Test Android, iOS, and Smart TV apps on real devices.",
  keywords: "mobile testing, device lab, on-premise testing, enterprise testing, Android testing, iOS testing, Appium, test automation",
  openGraph: {
    title: "RobusTest - Enterprise Mobile Device Lab Management",
    description: "Your Devices. Your Data. Your Control. Enterprise-grade mobile device lab management deployed on your premises.",
    type: "website",
    url: "https://www.robustest.com",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.variable} antialiased`}>
        <Header />
        <main>{children}</main>
        <Footer />
      </body>
    </html>
  );
}
