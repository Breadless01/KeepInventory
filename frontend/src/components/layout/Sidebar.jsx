import { NavLink } from 'react-router-dom';
import { useState } from "react";
import { ChevronRight, ChevronDown } from "lucide-react";
import "./sidebar.css";

export function Sidebar({ routes }) {
  const [openMenus, setOpenMenus] = useState({});

  function toggleMenu(id) {
    setOpenMenus((prev) => ({
      ...prev,
      [id]: !prev[id]
    }));
  }

  return (
    <aside className="ki-sidebar">
      <div className="ki-sidebar-logo">
        <span className="ki-sidebar-logo-icon">W</span>
        <span className="ki-sidebar-logo-text">KeepInventory</span>
      </div>

      <nav className="ki-sidebar-nav">
        {routes.map((route) => {
          if (route.children) {
            const isOpen = openMenus[route.id];
            const GroupIcon = route.icon;
            
            return (
              <div key={route.id} className="ki-nav-group">
                <a
                  className="ki-nav-item"
                  onClick={() => toggleMenu(route.id)}
                >
                  <span className="ki-nav-icon">
                    {GroupIcon && <GroupIcon size={18} strokeWidth={2} />}
                  </span>
                  <span className="ki-nav-label">{route.label}</span>
                  <span className="ki-nav-chevron">
                    {isOpen ? <ChevronDown size={18} strokeWidth={2}/> : <ChevronRight size={18} strokeWidth={2}/>}
                  </span>
                </a>

                <div
                  className={
                    "ki-nav-subitems " +
                    (isOpen ? "ki-nav-subitems--open" : "ki-nav-subitems--closed")
                  }
                >
                  {route.children.map((child) => {
                    const ChildIcon = child.icon;
                    return (
                      <NavLink
                        key={child.id}
                        to={child.path}
                        className={({ isActive }) =>
                          "ki-nav-item" +
                          (isActive ? " ki-nav-item--active" : "")
                        }
                      >
                        {ChildIcon && (
                          <span className="ki-nav-icon">
                            <ChildIcon size={16} strokeWidth={2} />
                          </span>
                        )}
                        <span>{child.label}</span>
                      </NavLink>
                    );
                  })}
                </div>
              </div>
            )
          }
          const Icon = route.icon;
          return (
            <NavLink
              key={route.id}
              to={route.path}
              end={route.path === "/"}
              className={({ isActive }) =>
                "ki-nav-item" + (isActive ? " ki-nav-item--active" : "")
              }
            >
              <span className="ki-nav-icon">
                {Icon && <Icon size={18} strokeWidth={2} />}
              </span>
              <span className="ki-nav-label">{route.label}</span>
            </NavLink>
          );
        })}
      </nav>
    </aside>
  );
}