"use client";
import { useEffect, useState } from "react";
import Sidebar from "./components/Sidebar";
type ResumoVendas = {
  clientes_ativos: number;
  quantidade_vendas: number;
  ticket_medio: number;
  total_vendas: number;
};
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
type Oportunidade = {
  id: number;
  cliente_id: number;
  titulo: string;
  etapa_funil: string;
  valor: number;
  data_criacao: string;
};
export default function Home() {
  const [resumo, setResumo] = useState<ResumoVendas | null>(null);
  const [clientes, setClientes] = useState<Cliente[]>([]);
  const [oportunidades, setOportunidades] = useState<Oportunidade[]>([]);
  const contarEtapa = (etapa: string) => {
    return oportunidades.filter(
      (oportunidade) => oportunidade.etapa_funil === etapa
    ).length;
  };
  useEffect(() => {
    fetch("http://localhost:8080/vendas/resumo")
      .then((resposta) => resposta.json())
      .then((dados: ResumoVendas) => setResumo(dados))
      .catch((erro) => console.error("Erro ao buscar resumo:", erro));
  }, []);
  useEffect(() => {
    fetch("http://localhost:8080/clientes")
      .then((resposta) => resposta.json())
      .then((dados: Cliente[]) => setClientes(dados))
      .catch((erro) => console.error("Erro ao buscar clientes:", erro));
  }, []);
  useEffect(() => {
    fetch("http://localhost:8080/funil")
      .then((resposta) => resposta.json())
      .then((dados: Oportunidade[]) => setOportunidades(dados))
      .catch((erro) => console.error("Erro ao buscar funil:", erro));
  }, []);
  return (
    <div className="flex min-h-screen bg-slate-100">
     <Sidebar />
     <main className="flex-1 p-8">
     <header className="mb-8">
  <h1 className="text-3xl font-bold text-slate-900">
    Visão Geral
  </h1>

  <p className="mt-2 text-slate-500">
    Bem-vindo ao painel de gerenciamento EnviePro.
  </p>
</header>
<section className="grid grid-cols-1 gap-6 md:grid-cols-3">
  <div className="rounded-xl bg-white p-6 shadow-sm">
    <p className="text-sm text-slate-500">Total de vendas</p>
    <h2 className="mt-2 text-2xl font-bold text-slate-900">
    {resumo
  ? resumo.total_vendas.toLocaleString("pt-BR", {
      style: "currency",
      currency: "BRL",
    })
  : "Carregando..."}
    </h2>
  </div>

  <div className="rounded-xl bg-white p-6 shadow-sm">
    <p className="text-sm text-slate-500">Quantidade de vendas</p>
    <h2 className="mt-2 text-2xl font-bold text-slate-900">
    {resumo ? resumo.quantidade_vendas : "Carregando..."}
    </h2>
  </div>

  <div className="rounded-xl bg-white p-6 shadow-sm">
    <p className="text-sm text-slate-500">Ticket médio</p>
    <h2 className="mt-2 text-2xl font-bold text-slate-900">
    {resumo
  ? resumo.ticket_medio.toLocaleString("pt-BR", {
      style: "currency",
      currency: "BRL",
    })
  : "Carregando..."}
    </h2>
  </div>
</section>
<section className="mt-8">
  <h2 className="mb-4 text-xl font-bold text-slate-900">
    Funil de vendas
  </h2>
  <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-5">
  <div className="rounded-xl bg-white p-4 shadow-sm">
    <h3 className="font-bold text-slate-800">Potencial cliente</h3>
    <p className="mt-2 text-2xl font-bold text-blue-600">{contarEtapa("Lead")}</p>
  </div>

  <div className="rounded-xl bg-white p-4 shadow-sm">
    <h3 className="font-bold text-slate-800">Contato</h3>
    <p className="mt-2 text-2xl font-bold text-blue-600">{contarEtapa("Contato")}</p>
  </div>

  <div className="rounded-xl bg-white p-4 shadow-sm">
    <h3 className="font-bold text-slate-800">Proposta</h3>
    <p className="mt-2 text-2xl font-bold text-blue-600">{contarEtapa("Proposta")}</p>
  </div>

  <div className="rounded-xl bg-white p-4 shadow-sm">
    <h3 className="font-bold text-slate-800">Negociação</h3>
    <p className="mt-2 text-2xl font-bold text-blue-600">{contarEtapa("Negociação")}</p>
  </div>

  <div className="rounded-xl bg-white p-4 shadow-sm">
    <h3 className="font-bold text-slate-800">Fechado</h3>
    <p className="mt-2 text-2xl font-bold text-blue-600">{contarEtapa("Fechado")}</p>
  </div>
</div>
</section>
<section className="mt-8">
  <h2 className="mb-4 text-xl font-bold text-slate-900">
    Clientes recentes
  </h2>
  <div className="overflow-x-auto rounded-xl bg-white shadow-sm">
  <table className="w-full text-left">
    <thead className="border-b border-slate-200 bg-slate-50">
      <tr>
        <th className="p-4 text-sm font-semibold text-slate-600">Nome</th>
        <th className="p-4 text-sm font-semibold text-slate-600">Empresa</th>
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
    <td className="p-4">
      <span className="rounded-full bg-green-100 px-3 py-1 text-xs font-medium text-green-700">
        {cliente.status}
      </span>
    </td>
    <td className="p-4 text-sm text-slate-700">{cliente.responsavel}</td>
    <td className="p-4 text-sm text-slate-700">
      {new Date(cliente.ultimo_contato).toLocaleDateString("pt-BR")}
    </td>
  </tr>
))}
    </tbody>
  </table>
</div>
</section>
</main>
    </div>
  );
}