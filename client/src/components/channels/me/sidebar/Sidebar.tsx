"use client";

import DMList from "./DMList";
import SidebarBottom from "./SidebarBottom";
import SidebarTop from "./SidebarTop";

export default function Sidebar() {
  return (
    <div className="flex flex-col h-full">
      <SidebarTop />
      <DMList />
      <SidebarBottom />
    </div>
  );
}
