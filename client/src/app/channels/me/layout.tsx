import DMList from "@/components/channels/me/DMList";

export default function MeLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex grow shrink">
      <div className="grow-0 shrink-0 w-[240px] p-3">
        <DMList />
      </div>
      <div className="bg-neutral-600 grow shrink">{children}</div>
    </div>
  );
}
