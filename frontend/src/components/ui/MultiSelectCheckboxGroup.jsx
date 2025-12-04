import "./MultiSelectCheckboxGroup.css"
import {useState} from "react";
import { ChevronRight, ChevronDown } from "lucide-react";

export default function MultiSelectCheckboxGroup({ label, options, values, onChange, useKey = true }) {
    const [open, setOpen] = useState(false);

    function toggle(id) {
        if (values.includes(id)) {
            onChange(values.filter((v) => v !== id));
        } else {
            onChange([...values, id]);
        }
    }

    const selectedCount = values?.length ?? 0;

    return (
      <div className="ki-facet">
        {label && <div className="ki-facet-label">{label}</div>}
        <div
          className={`ki-ms-control ${open ? "ki-ms-control--open" : ""}`}
          onClick={() => setOpen((o) => !o)}
          onBlur={() => setOpen((o) => !o)}
        >
          <span className="ki-ms-value">
            {selectedCount === 0 && <span className="ki-ms-placeholder">Keine Auswahl</span>}
            {selectedCount === 1 && <span>{selectedCount} Wert ausgewählt</span>}
            {selectedCount > 1 && <span>{selectedCount} Werte ausgewählt</span>}
          </span>
          <span className="ki-ms-chevron">{open ? <ChevronDown size={18} strokeWidth={2}/> : <ChevronRight size={18} strokeWidth={2}/>}</span>
        </div>
        {open && (
          <div className="ki-ms-dropdown">
            {options.map((opt) => {
              const id = opt.ID ?? opt.id;
              const name = opt.Name ?? opt.name;
              const count = opt.Count ?? opt.count;
              const isSelected = useKey ? values.includes(id) : values.includes(name);

              return (
                <div
                  key={id}
                  className={
                    "ki-ms-option" + (isSelected ? " ki-ms-option--selected" : "")
                  }
                  onClick={(e) => {
                    e.stopPropagation();
                    if (useKey) toggle(id);
                    else toggle(name)
                  }}
                >
                  <span className="ki-ms-option-name">{name}</span>
                  {typeof count === "number" && (
                    <span className="ki-ms-option-count">{count}</span>
                  )}
                </div>
              )
            })}
          </div>
        )}
      </div>
    );
}