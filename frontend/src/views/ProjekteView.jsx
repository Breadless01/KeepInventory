import { useEffect, useState } from "react";
import { FlexTable } from "../components/ui/FlexTable.jsx";
import {NewProjektModal} from "../components/special/NewProjektModal.jsx";
import FacetFilterPanel from "../components/ui/FacetFilterPanel.jsx"
import { Plus } from "lucide-react";
import {
  FilterProjekte,
  GetFilterConfig,
} from "../../wailsjs/go/backend/App";
import { useToasts } from "../components/ui/ToastContainer.jsx";
import { Searchbar } from "../components/ui/Searchbar.jsx";

export default function ProjekteView() {
  const [projekte, setProjekte] = useState([]);
  const [modalOpen, setModalOpen] = useState(false);

  const columns = [
    { id: "id", label: "ID", field: "ID", width: 0.5, align: "center" },
    { id: "name", label: "Name", field: "Name", width: 2 },
    { id: "kunde", label: "Kunde", field: "Kunde", width: 2 },
  ];

  // Filter-spezifischer State
  const [filterConfig, setFilterConfig] = useState(null);
  const [filterState, setFilterState] = useState({});
  const [filterIds, setFilterIds] = useState([]);
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
      const res = await FilterProjekte({
        page: pageNumber,
        pageSize,
        filters: state,
      });

      setProjekte(res.items || []);
      setFacets(res.facets || {});
      setTotal(res.total || 0);
      setPage(pageNumber);
    } catch (e) {
      addToast({
        title: "Fehler beim Laden der Projekte",
        message: String(e),
        type: "error",
        mode: "static",
      });
      console.error(e);
    }
  }

  useEffect(() => {
    setFilterState({ ...filterState, ["id"]: filterIds })
  }, [filterIds]);

  useEffect(() => {
    if (!filterConfig) return;
    applyFilter(filterConfig, filterState, 1);
  }, [filterState, filterConfig]);

  function handleProjektCreated() {
    if (!filterConfig) {
      applyFilter({ resources: [] }, {}, 1);
    } else {
      applyFilter(filterConfig, filterState, page);
    }
  }

  const safeProjekte = projekte || [];

  let facetFieldConfigs = [];
  if (filterConfig && Array.isArray(filterConfig.resources)) {
    const projektRes = filterConfig.resources.find(
      (r) => r.resource === "projekte"
    );
    if (projektRes && Array.isArray(projektRes.fields)) {
      facetFieldConfigs = projektRes.fields.filter((f) => f.enabled);
    }
  }
  return (
    <div className="ki-content">
      <Searchbar
        objType="projekt"
        onEnter={setFilterIds}
      />
      <div className="ki-card">
        <div className="ki-header-row">
          <h2 className="ki-card-title">Projekte</h2>
          <button className="ki-add-btn" title="Neues Projekt anlegen" onClick={() => setModalOpen(true)}>
            <Plus size={16} strokeWidth={4} />
          </button>
        </div>
        {facetFieldConfigs.length > 0 && (
          <div style={{ marginBottom: "1rem" }}>
            <FacetFilterPanel
              total={total}
              useKeyValues={false}
              facets={facets}
              filterState={filterState}
              onChange={setFilterState}
              fieldConfigs={facetFieldConfigs}
            />
          </div>
        )}
        <FlexTable columns={columns} data={safeProjekte} />
      </div>
      <NewProjektModal
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        onCreated={() => {
          handleProjektCreated();
        }}
      />
    </div>
  );
}
