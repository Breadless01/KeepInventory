import { Routes, Route, useLocation } from "react-router-dom";
import { Sidebar } from "./components/layout/Sidebar.jsx";
import { routes } from "./routes.js";
import { ToastProvider } from "./components/ui/ToastContainer.jsx";

function App() {
  const location = useLocation();
  const activeRoute =
    routes.find((r) => r.path === location.pathname) || routes[0];

  return (
    <ToastProvider>
      <div className="ki-shell">
        <Sidebar routes={routes} />

        <main className="ki-main">
          <header className="ki-header">
            <h1 className="ki-header-title">{activeRoute.label}</h1>
          </header>

          <Routes>
            {routes.map((r) => {
              if (r.children) {
                return r.children.map((child) => {
                  const Component = child.component;
                  return (
                    <Route
                    key={child.id}
                    path={child.path}
                    element={<Component />}
                    />
                  );
                });
              }
              
              const Component = r.component;
              return (
                <Route
                key={r.id}
                path={r.path}
                element={<Component />}
                />
              );
            })}
          </Routes>
        </main>
      </div>
    </ToastProvider>
  );
}

export default App;
