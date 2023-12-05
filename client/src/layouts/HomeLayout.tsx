import LeftNavBar from "@/components/channels/LeftNavBar";

export default function HomeLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="h-screen flex bg-neutral-700">
      <LeftNavBar />
      {children}
    </div>
  );
}
