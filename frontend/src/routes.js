import InventoryView from "./views/InventoryView.jsx";
import KundenView from "./views/KundenView.jsx";
import ProjekteView from "./views/ProjekteView.jsx";
import FilterSettingsView from "./views/FilterSettingsView.jsx";

import { Database, Users, Folder, Settings, Diamond, Funnel, Truck } from "lucide-react";
import {createRef} from "react";
import HerstellerView from "./views/HerstellerView.jsx";
import LieferantenView from "./views/HerstellerView.jsx";

export const routes = [
  {
    id: "db",
    label: "DB",
    icon: Database,
    children: [
      {
        path: "/db/bauteile",
        id: "bauteile",
        nodeRef: createRef(),
        label: "Bauteile",
        icon: Diamond,
        component: InventoryView,
      },
      {
        path: "/db/kunden",
        id: "kunden",
        nodeRef: createRef(),
        label: "Kunden",
        icon: Users,
        component: KundenView
      },
      {
        path: "/db/projekte",
        id: "projekte",
        nodeRef: createRef(),
        label: "Projekte",
        icon: Folder,
        component: ProjekteView
      },
      {
        path: "/db/lieferanten",
        id: "lieferanten",
        nodeRef: createRef(),
        label: "Lieferanten",
        icon: Truck,
        component: LieferantenView
      }
    ]
  },
  {
    id: "settings",
    label: "Settings",
    icon: Settings,
    children: [
        {
            path: "/settings/filterSettings",
            id: "filterSettings",
            nodeRef: createRef(),
            label: "Filter",
            icon: Funnel,
            component: FilterSettingsView
        }
    ],
  }
];