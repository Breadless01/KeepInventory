import "./flexTable.css";
import { useEffect, useMemo, useState } from "react";
import { ChevronRight, ChevronLeft } from "lucide-react";
import { Modal } from "../ui/Modal.jsx";


export function FlexTable({ columns, data, pageSize = 10, onNext, onPrevious}) {
  const [currentPage, setCurrentPage] = useState(1);
  const [showModal, setShowModal] = useState(false);
  const [selected, setSelected] = useState(null);
  const [saving, setSaving] = useState(false);

  const safeData = data || [];
  const totalItems = safeData.length;
  const totalPages = Math.max(1, Math.ceil(totalItems / pageSize));

  useEffect(() => {
    if (currentPage > totalPages) {
      setCurrentPage(totalPages);
    }
  }, [totalPages, currentPage]);

  const pageData = useMemo(() => {
    const start = (currentPage - 1) * pageSize;
    return safeData.slice(start, start + pageSize);
  }, [safeData, currentPage, pageSize]);

  const startIndex = totalItems === 0 ? 0 : (currentPage - 1) * pageSize + 1;
  const endIndex = totalItems === 0
    ? 0
    : Math.min(currentPage * pageSize, totalItems);

  if (!columns || columns.length === 0) {
    return <div className="ki-dt-empty">Keine Spalten definiert.</div>;
  }

  const handleModalOpen = (row) => {
    setSelected(row);
    setShowModal(true);
  }

  function handleModalClose() {
    setShowModal(false);
  }

  return (
    <div className="ki-dt-card">
      {/* Header */}
      <div className="ki-dt-header">
        {columns.map((col) => (
          <div
            key={col.id}
            className="ki-dt-cell ki-dt-cell--head"
            style={{
              flex: col.width ?? 1,
              justifyContent: mapAlign(col.align),
            }}
          >
            {col.label}
          </div>
        ))}
      </div>

      {/* Rows */}
      <div className="ki-dt-body">
        {(pageData.length === 0) ? (
          <div className="ki-dt-row ki-dt-row--empty">
            <div className="ki-dt-cell">Keine Einträge vorhanden.</div>
          </div>
        ) : (
          pageData.map((row, rowIndex) => (
            <div
              key={row.id ?? row.ID ?? rowIndex}
              className="ki-dt-row"
              onDoubleClick={() => handleModalOpen(row)}
            >
              {columns.map((col) => (
                <div
                  key={col.id}
                  className="ki-dt-cell"
                  style={{
                    flex: col.width ?? 1,
                    justifyContent: mapAlign(col.align),
                  }}
                >
                  {col.render
                    ? col.render(row)
                    : col.field
                    ? row[col.field]
                    : null}
                </div>
              ))}
            </div>
          ))
        )}
      </div>
      {/* Pagination */}
      <div className="ki-dt-footer">
        <div className="ki-dt-page-info">
          {totalItems === 0
            ? "Keine Einträge"
            : `Zeige ${startIndex}–${endIndex} von ${totalItems}`}
        </div>

        <div className="ki-dt-pagination">
          <button
            type="button"
            className="ki-dt-page-btn"
            onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}
            disabled={currentPage === 1}
          >
            <ChevronLeft size={18} strokeWidth={2}/>
          </button>
          <span className="ki-dt-page-indicator">
            Seite {currentPage} / {totalPages}
          </span>
          <button
            type="button"
            className="ki-dt-page-btn"
            onClick={() =>
              setCurrentPage((p) => Math.min(totalPages, p + 1))
            }
            disabled={currentPage === totalPages || totalItems === 0}
          >
            <ChevronRight size={18} strokeWidth={2}/>
          </button>
        </div>
      </div>
      {showModal && selected ?
        <Modal
          title={"Bauteil: " + selected[columns.find(obj => obj.id === "name").field]}
          onClose={handleModalClose}
        >
          <form className="ki-form" onSubmit={() => {
            console.log(selected)
          }}>
            {columns.map((column) => {
              if (!["id", "erstelldatum", "sachnummer"].includes(column.id)) {
                return(
                  <div className="ki-form-group" key={column.id}>
                    <label>{column.label}</label>
                    <input
                      className="ki-input"
                      value={selected[columns.find(obj => obj.id === column.id).label]}
                      onChange={(e) => {
                        let obj = {...selected}
                        obj[column.field] = e.target.value;
                        setSelected(obj);
                      }}
                    />
                  </div>
                )
              }
            })}
            <div className="ki-form-actions">
              <button
                type="button"
                className="ki-btn-secondary"
                onClick={handleModalClose}
                disabled={saving}
              >
                Abbrechen
              </button>
              <button
                type="submit"
                className="ki-btn-primary"
                disabled={saving}
              >
                {saving ? "Speichern…" : "Bauteil anlegen"}
              </button>
            </div>
          </form>
        </Modal>
      : ""}
    </div>
  );
}

function mapAlign(align) {
  switch (align) {
    case "center":
      return "center";
    case "right":
      return "flex-end";
    case "left":
    default:
      return "flex-start";
  }
}
