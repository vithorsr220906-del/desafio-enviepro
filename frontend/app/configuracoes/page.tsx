
import Sidebar from "../components/Sidebar";

export default function Configuracoes() {
  return (
    <div className="flex min-h-screen bg-slate-100">
      <Sidebar />

      <main className="flex-1 p-8">
        <h1 className="text-3xl font-bold text-slate-900">
          Configurações
        </h1>

        <p className="mt-2 text-slate-500">
          Gerencie as preferências do sistema.
        </p>

        <div className="mt-8 max-w-3xl space-y-6">
          <section className="rounded-xl bg-white p-6 shadow-sm">
            <h2 className="text-lg font-semibold text-slate-800">
              Informações da empresa
            </h2>

            <p className="mt-2 text-sm text-slate-500">
              EnviePro
            </p>
            <p className="mt-1 text-sm text-slate-500">
              Sistema de gerenciamento de clientes e vendas.
            </p>
          </section>

          <section className="rounded-xl bg-white p-6 shadow-sm">
            <h2 className="text-lg font-semibold text-slate-800">
              Preferências
            </h2>

            <div className="mt-4 flex items-center justify-between border-b border-slate-100 py-3">
              <div>
                <p className="font-medium text-slate-700">
                  Idioma
                </p>
                <p className="text-sm text-slate-500">
                  Idioma utilizado no sistema
                </p>
              </div>

              <span className="text-sm text-slate-600">
                Português (Brasil)
              </span>
            </div>

            <div className="flex items-center justify-between py-3">
              <div>
                <p className="font-medium text-slate-700">
                  Versão
                </p>
                <p className="text-sm text-slate-500">
                  Versão atual do sistema
                </p>
              </div>

              <span className="text-sm text-slate-600">
                1.0.0
              </span>
            </div>
          </section>
        </div>
      </main>
    </div>
  );
}
