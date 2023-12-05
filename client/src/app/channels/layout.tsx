import HomeLayout from "@/layouts/HomeLayout";

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <HomeLayout>{children}</HomeLayout>;
}
