import {useEffect, useState} from "react";
import { Search } from "../../../wailsjs/go/backend/App.js"
import "./Searchbar.css"

export function Searchbar({objType, onEnter}) {
  const [filterIds, setFilterIds] = useState([]);

  const [searchTerm, setSearchTerm] = useState("");
  const [suggestions, setSuggestions] = useState([]);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [isSearching, setIsSearching] = useState(false);

  useEffect(() => {
    if (!searchTerm || searchTerm.trim().length < 2) {
      setSuggestions([]);
      setShowSuggestions(false);
      return;
    }

    setIsSearching(true);
    const handle = setTimeout(async () => {
      try {
        const res = await Search(searchTerm, objType, 10);
        setSuggestions(res || []);
        setShowSuggestions(true);
      } catch (err) {
        console.error("Search error:", err);
      } finally {
        setIsSearching(false);
      }
    }, 200);

    return () => clearTimeout(handle);
  }, [searchTerm]);

  useEffect(() => {
    setFilterIds(suggestions.map((suggestion) => {
      return (
        suggestion.id
      )
    }))
  }, [suggestions])

  return (
    <div className="ki-search-wrapper">
      <input
        type="text"
        className="ki-input ki-input-search"
        placeholder="Suchen…"
        value={searchTerm}
        onKeyDown={(type) => {
          if (type.code === "Enter") {
            onEnter(filterIds)
          }
        }}
        onChange={(e) => setSearchTerm(e.target.value)}
        onFocus={() => {
          if (suggestions.length > 0) setShowSuggestions(true);
        }}
        onBlur={() => {
          // kleines Delay, damit Klick auf Suggestion noch geht
          setTimeout(() => setShowSuggestions(false), 150);
        }}
      />
      {isSearching && <div className="ki-search-spinner" />}

      {showSuggestions && suggestions.length > 0 && (
        <div className="ki-search-suggestions">
          {suggestions.map((s) => (
            <button
              key={s.id}
              type="button"
              className="ki-search-suggestion"
              onMouseDown={(e) => e.preventDefault()}
              onClick={() => {
                onEnter([s.id])
                setShowSuggestions(false);
              }}
            >
              <div className="ki-search-suggestion-main">
                      <span className="ki-search-suggestion-title">
                        {s.label}
                      </span>
                {s.subtitle && (
                  <span className="ki-search-suggestion-sub">
                          {s.subtitle}
                        </span>
                )}
              </div>
            </button>
          ))}
        </div>
      )}
    </div>
  )
}