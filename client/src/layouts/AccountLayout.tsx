export default function AccountLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-screen min-w-screen items-center justify-center">
      {children}
    </div>
  );
}
