// Package ask provides utility functions for conditional operations and zero value checking.
//
// This package offers ternary-like operations and null coalescing functionality using Go generics,
// designed to reduce boilerplate code in conditional assignments and provide a more expressive
// way to handle default values and conditional logic.
//
// # Key Features
//
//   - High-performance conditional operations with type safety
//   - Zero-dependency implementation using only standard library
//   - Optimized for common types to avoid reflection overhead
//   - Comprehensive zero value and empty value checking
//   - SQL-like COALESCE functionality for multiple value fallbacks
//
// # Basic Usage
//
//	// Ternary operator
//	result := ask.If(condition, "true value", "false value")
//
//	// Null coalescing
//	value := ask.Ifelse(maybeEmpty, "default value")
//
//	// Multiple value coalescing
//	final := ask.Coalesce(first, second, third, "fallback")
//
// # Performance
//
// The package is optimized for performance with type-specific fast paths for common types
// (bool, int, string, etc.) to avoid reflection overhead. Benchmark tests show minimal
// performance impact compared to native Go conditionals.
//
// # Type Support
//
// All functions work with any Go type through generics, with special handling for:
//   - Basic types (bool, int, float, string, etc.)
//   - Pointer types (checked for nil)
//   - Slice, map, and channel types (checked for nil and length)
//   - Interface types including error
//   - Struct types and arrays
//
// # Examples
//
// See the examples directory for comprehensive usage examples including:
//   - Basic conditional operations
//   - Configuration management
//   - API response building
//   - Template data preparation
//   - Error handling patterns
package ask