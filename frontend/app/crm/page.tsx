"use client";

import { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";

type Cliente = {
  id: number;
  nome: string;
  empresa: string;
  email: string;
  telefone: string;
  status: string;
  responsavel: string;
  ultimo_contato: string;
};

export default function CRM() {
  const [clientes, setClientes] = useState<Cliente[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/clientes")
      .then((resposta) => resposta.json())
      .then((dados: Cliente[]) => setClientes(dados))
      .catch((erro) => console.error("Erro ao buscar clientes:", erro));
  }, []);

  return (
    <div className="flex min-h-screen bg-slate-100">
      <Sidebar />

      <main className="flex-1 p-8">
        <h1 className="text-3xl font-bold text-slate-900">
          Gerenciamento de Clientes
        </h1>

        <p className="mt-2 text-slate-500">
          Consulte e acompanhe seus clientes cadastrados.
        </p>
        <div className="mt-8 overflow-x-auto rounded-xl bg-white shadow-sm">
  <table className="w-full text-left">
    <thead className="border-b border-slate-200 bg-slate-50">
      <tr>
        <th className="p-4 text-sm font-semibold text-slate-600">Nome</th>
        <th className="p-4 text-sm font-semibold text-slate-600">Empresa</th>
        <th className="p-4 text-sm font-semibold text-slate-600">E-mail</th>
        <th className="p-4 text-sm font-semibold text-slate-600">Telefone</th>
        <th className="p-4 text-sm font-semibold text-slate-600">Status</th>
        <th className="p-4 text-sm font-semibold text-slate-600">Responsável</th>
        <th className="p-4 text-sm font-semibold text-slate-600">Último contato</th>
      </tr>
    </thead>

    <tbody>
      {clientes.map((cliente) => (
        <tr key={cliente.id} className="border-b border-slate-100">
          <td className="p-4 text-sm text-slate-700">{cliente.nome}</td>
          <td className="p-4 text-sm text-slate-700">{cliente.empresa}</td>
          <td className="p-4 text-sm text-slate-700">{cliente.email}</td>
          <td className="p-4 text-sm text-slate-700">{cliente.telefone}</td>
          <td className="p-4 text-sm text-slate-700">{cliente.status}</td>
          <td className="p-4 text-sm text-slate-700">{cliente.responsavel}</td>
          <td className="p-4 text-sm text-slate-700">
            {new Date(cliente.ultimo_contato).toLocaleDateString("pt-BR")}
          </td>
        </tr>
      ))}
    </tbody>
  </table>
</div>
      </main>
    </div>
  );
}