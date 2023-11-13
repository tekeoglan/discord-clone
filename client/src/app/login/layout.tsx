import AccountLayout from "@/layouts/AccountLayout";

export default function LoginLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <AccountLayout>{children}</AccountLayout>;
}
