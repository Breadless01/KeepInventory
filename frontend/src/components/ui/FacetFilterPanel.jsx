import MultiSelectCheckboxGroup from "./MultiSelectCheckboxGroup";
import "./FacetFilterPanel.css"
import {useEffect, useState} from "react";
import { ChevronRight, ChevronDown } from "lucide-react";
import { ListLieferanten } from "../../../wailsjs/go/backend/App";

export default function FacetFilterPanel({ facets, filterState, onChange, fieldConfigs, useKeyValues = true, total }) {
  const [open, setOpen] = useState(false);
  return (
    <div className={"ki-facet-wrapper " + (open ? "ki-facet-wrapper-open" : "")}>
      <div className="ki-header-row ki-filter-header" onClick={() => setOpen(!open)}>
        <h3 className="ki-filter-title">Filter {total ? "("+total+")" : ""}</h3>

        {open ? <ChevronDown size={18} strokeWidth={2}/> : <ChevronRight size={18} strokeWidth={2}/>}
      </div>
      <div className="ki-facet-panel">
      {fieldConfigs
        .filter((f) => f.enabled)
        .map((f) => {
          const field = f.field;
          const label = f.label;
          const options = facets[field] || [];
          const values = filterState[field] || [];
          return (
            <MultiSelectCheckboxGroup
              key={field}
              useKey={useKeyValues}
              label={label}
              options={options}
              values={values}
              onChange={(newValues) =>
                onChange({ ...filterState, [field]: newValues })
              }
            />
          );
        })}
      </div>
    </div>
  );
}
