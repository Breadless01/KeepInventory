import { useEffect, useState } from "react";
import { FlexTable } from "../components/ui/FlexTable.jsx";
import { NewKundeModal } from "../components/special/NewKundeModal.jsx";
import FacetFilterPanel from "../components/ui/FacetFilterPanel.jsx"
import { Plus } from "lucide-react";
import {
  FilterKunden,
  GetFilterConfig
} from "../../wailsjs/go/backend/App";
import { useToasts } from "../components/ui/ToastContainer.jsx";

export default function KundenView() {
  const [kunden, setKunden] = useState([]);
  const [modalOpen, setModalOpen] = useState(false);

  const columns = [
    { id: "id", label: "ID", field: "ID", width: 0.5, align: "center" },
    { id: "name", label: "Name", field: "Name", width: 2 },
    { id: "sitz", label: "Sitz", field: "Sitz", width: 2 },
  ];

  // Filter-spezifischer State
  const [filterConfig, setFilterConfig] = useState(null);
  const [filterState, setFilterState] = useState({});
  const [facets, setFacets] = useState({});
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const pageSize = 50;

  const { addToast } = useToasts();

  useEffect(() => {
    async function init() {
      try {
        const cfg = await GetFilterConfig();
        setFilterConfig(cfg);

        await applyFilter(cfg, {}, 1);
      } catch (e) {
        addToast({
          title: "Fehler beim Laden der Filter-Konfiguration",
          message: String(e),
          type: "error",
          mode: "static",
        });
        console.error(e);
      }
    }

    init();
  }, [])

  async function applyFilter(cfg, state, pageNumber) {
    try {
      const res = await FilterKunden({
        page: pageNumber,
        pageSize,
        filters: state,
      });

      setKunden(res.items || []);
      setFacets(res.facets || {});
      setTotal(res.total || 0);
      setPage(pageNumber);
    } catch (e) {
      addToast({
        title: "Fehler beim Laden der Kunden",
        message: String(e),
        type: "error",
        mode: "static",
      });
      console.error(e);
    }
  }

  useEffect(() => {
    if (!filterConfig) return;
    applyFilter(filterConfig, filterState, 1);
  }, [filterState, filterConfig]);

  function handleKundeCreated() {
    if (!filterConfig) {
      applyFilter({ resources: [] }, {}, 1);
    } else {
      applyFilter(filterConfig, filterState, page);
    }
  }

  const safeKunden = kunden || [];

  let facetFieldConfigs = [];
  if (filterConfig && Array.isArray(filterConfig.resources)) {
    const kundeRes = filterConfig.resources.find(
      (r) => r.resource === "kunden"
    );
    if (kundeRes && Array.isArray(kundeRes.fields)) {
      facetFieldConfigs = kundeRes.fields.filter((f) => f.enabled);
    }
  }
  return (
    <div className="ki-content">
      <div className="ki-card">
        <div className="ki-header-row">
          <h2 className="ki-card-title">Kunden</h2>
          <button className="ki-add-btn" title="Neuen Kunden anlegen" onClick={() => setModalOpen(true)}>
            <Plus size={16} strokeWidth={4} />
          </button>
        </div>
        {facetFieldConfigs.length > 0 && (
          <div style={{ marginBottom: "1rem" }}>
            <FacetFilterPanel
              useKeyValues={false}
              facets={facets}
              filterState={filterState}
              onChange={setFilterState}
              fieldConfigs={facetFieldConfigs}
            />
          </div>
        )}
        <FlexTable columns={columns} data={safeKunden} />
      </div>
      <NewKundeModal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        onCreated={() => {
          handleKundeCreated();
        }}
      />
    </div>
  );
}
