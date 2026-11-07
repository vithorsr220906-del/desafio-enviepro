"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import {
  Home,
  Users,
  Settings,
  LogOut,
  ChevronLeft,
  ChevronRight,
} from "lucide-react";
  import { useState } from "react";

  export default function Sidebar() {
    const pathname = usePathname();
    const [aberta, setAberta] = useState(true);
    return (
        <aside
  className={`flex h-screen ${aberta ? "w-64" : "w-20"} flex-col bg-slate-900 p-4 text-white transition-all duration-300`}
>
        <div className="mb-8 flex items-center gap-3">
  <div className="rounded-lg bg-blue-600 p-2">
    <Users size={24} />
  </div>

  {aberta && (
  <span className="text-xl font-bold">
    EnviePro
  </span>
)}
</div>
<button
  onClick={() => setAberta(!aberta)}
  className="mb-4 flex items-center justify-center rounded-lg p-2 hover:bg-slate-800"
>
  {aberta ? <ChevronLeft size={20} /> : <ChevronRight size={20} />}
</button>
<nav className="flex flex-col gap-2">
<Link
  href="/"
  className={`flex items-center gap-3 rounded-lg px-3 py-3 text-left ${
    pathname === "/" ? "bg-blue-600" : "hover:bg-slate-800"
  }`}
>
  <Home size={20} />
  {aberta && "Início"}
</Link>

<Link
  href="/crm"
  className={`flex items-center gap-3 rounded-lg px-3 py-3 text-left ${
    pathname === "/crm" ? "bg-blue-600" : "hover:bg-slate-800"
  }`}
>
  <Users size={20} />
  {aberta && "CRM"}
</Link>

<Link
  href="/configuracoes"
  className={`flex items-center gap-3 rounded-lg px-3 py-3 text-left ${
    pathname === "/configuracoes" ? "bg-blue-600" : "hover:bg-slate-800"
  }`}
>
  <Settings size={20} />
  {aberta && "Configurações"}
</Link>
</nav>
<div className="mt-auto border-t border-slate-700 pt-4">
  <button className="flex w-full items-center gap-3 rounded-lg px-3 py-3 text-left hover:bg-slate-800">
  <LogOut size={20} />
  {aberta && "Sair"}
  </button>
</div>
        </aside>
      );
  }