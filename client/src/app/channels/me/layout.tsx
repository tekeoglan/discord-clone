import Sidebar from "@/components/channels/me/sidebar/Sidebar";

export default function MeLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="relative flex justify-start items-stretch grow shrink">
      <div className="grow-0 shrink-0 w-[240px]">
        <Sidebar />
      </div>
      {children}
    </div>
  );
}
