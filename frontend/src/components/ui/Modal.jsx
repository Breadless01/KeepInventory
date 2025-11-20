import { createPortal } from "react-dom";
import "./modal.css";

export function Modal({ title, children, onClose }) {
  return createPortal(
    <div className="ki-modal-overlay" onMouseDown={onClose}>
      <div
        className="ki-modal"
        onMouseDown={(e) => e.stopPropagation()} // verhindert Schließen bei Klick ins Modal
      >
        <div className="ki-modal-header">
          <h3>{title}</h3>
          <button className="ki-modal-close" onClick={onClose}>
            ×
          </button>
        </div>

        <div className="ki-modal-body">{children}</div>
      </div>
    </div>,
    document.body
  );
}
