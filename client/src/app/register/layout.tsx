import AccountLayout from "@/layouts/AccountLayout";

export default function RegisterLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <AccountLayout>{children}</AccountLayout>;
}
