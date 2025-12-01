import MultiSelectCheckboxGroup from "./MultiSelectCheckboxGroup";
import "./FacetFilterPanel.css"

export default function FacetFilterPanel({ facets, filterState, onChange, fieldConfigs, useKeyValues = true }) {
  return (
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
  );
}
