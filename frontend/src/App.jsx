import { Routes, Route, useLocation } from "react-router-dom";
import {CSSTransition, SwitchTransition, TransitionGroup} from "react-transition-group"
import { Sidebar } from "./components/layout/Sidebar.jsx";
import { routes } from "./routes.js";
import LandingView from "./views/Landing.jsx";
import { ToastProvider } from "./components/ui/ToastContainer.jsx";
import {createRef} from "react";

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
            <TransitionGroup>
              <CSSTransition
                timeout={300}
                classNames='fade'
                key={location.key}
              >
                <Routes location={location}>
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
                  <Route path="/" element={<LandingView/>}/>
                </Routes>
              </CSSTransition>
            </TransitionGroup>
        </main>
      </div>
    </ToastProvider>
  );
}

export default App;
