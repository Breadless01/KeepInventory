import InventoryView from "./views/InventoryView.jsx";
import KundenView from "./views/KundenView.jsx";
import ProjekteView from "./views/ProjekteView.jsx";
import SettingsView from "./views/SettingsView.jsx";

import { Database, Users, Folder, Settings, Diamond } from "lucide-react";

export const routes = [
  {
    id: "db",
    label: "DB",
    icon: Database,
    children: [
      {
        path: "/bauteile",
        id: "bauteile",
        label: "Bauteile",
        icon: Diamond,
        component: InventoryView,
      },
      {
        path: "/kunden",
        id: "kunden",
        label: "Kunden",
        icon: Users,
        component: KundenView
      },
      {
        path: "/projekte",
        id: "projekte",
        label: "Projekte",
        icon: Folder,
        component: ProjekteView
      }
    ]
  },
  {
    path: "/",
    id: "settings",
    label: "Settings",
    icon: Settings,
    component: SettingsView,
  }
];