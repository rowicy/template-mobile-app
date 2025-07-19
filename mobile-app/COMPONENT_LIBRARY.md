# Gluestack-UI Component Library Setup

This project has been configured with Gluestack-UI v2, a comprehensive React Native component library that uses Tailwind CSS for styling.

## What was added

### 1. Component Library
- **Gluestack-UI v2**: A production-ready component library
- **NativeWind**: Tailwind CSS for React Native
- **Comprehensive components**: Button, Text, Input, Card, Heading, and more

### 2. Configuration Files
- `tailwind.config.js`: Tailwind CSS configuration with Gluestack theme
- `global.css`: Base Tailwind CSS imports
- `gluestack-ui.config.json`: Gluestack-UI configuration
- Updated `babel.config.js` and `metro.config.js` for proper bundling

### 3. Provider Setup
- `components/ui/gluestack-ui-provider/`: Theme provider with light/dark mode support
- Integrated with existing app layout in `app/_layout.tsx`

## Available Components

The following UI components are now available:

### Text & Typography
```tsx
import { Text } from "@/components/ui/text";
import { Heading } from "@/components/ui/heading";

<Heading size="xl">Page Title</Heading>
<Text size="md" className="text-typography-700">
  Body text with custom styling
</Text>
```

### Buttons
```tsx
import { Button, ButtonText } from "@/components/ui/button";

<Button action="primary" variant="solid" size="md">
  <ButtonText>Click me</ButtonText>
</Button>
```

### Cards
```tsx
import { Card } from "@/components/ui/card";

<Card className="p-4 mb-4">
  <Text>Card content</Text>
</Card>
```

### Input Fields
```tsx
import { Input, InputField } from "@/components/ui/input";

<Input>
  <InputField placeholder="Enter text..." />
</Input>
```

## Styling with Tailwind CSS

All components support Tailwind CSS classes through the `className` prop:

```tsx
<Text className="text-center text-lg font-bold text-primary-600">
  Styled text
</Text>

<Card className="bg-white shadow-lg rounded-lg p-6 mx-4">
  Card with custom styling
</Card>
```

## Theme System

The setup includes a comprehensive color system:
- **Primary**: Main brand colors
- **Secondary**: Secondary brand colors  
- **Success**: Green tones for success states
- **Error**: Red tones for error states
- **Warning**: Orange/yellow tones for warnings
- **Info**: Blue tones for informational content

Colors are available in shades from 50-950:
```tsx
<Text className="text-primary-500">Primary text</Text>
<View className="bg-success-100">Success background</View>
```

## Adding More Components

To add additional Gluestack-UI components:

```bash
npx gluestack-ui add [component-name]
```

Example:
```bash
npx gluestack-ui add checkbox radio badge
```

## Benefits

### ✅ Well Maintained
- Active development and regular updates
- Strong community support
- Enterprise-ready

### ✅ Tailwind-based
- Familiar CSS utility classes
- Responsive design capabilities
- Easy customization

### ✅ Rich Component Library
- 50+ production-ready components
- Consistent design system
- Accessibility built-in

### ✅ Easy Customization
- Theme customization through `tailwind.config.js`
- Component-level styling with `className`
- No migration needed when requirements change

## Live Demo

Check the "Components" tab in the app to see working examples of all integrated components.

## Migration from Existing Components

You can gradually migrate from existing styled components to Gluestack-UI:

### Before (StyleSheet)
```tsx
const styles = StyleSheet.create({
  button: {
    backgroundColor: '#0066cc',
    padding: 12,
    borderRadius: 8,
  },
  text: {
    color: 'white',
    fontSize: 16,
    fontWeight: 'bold',
  },
});

<Pressable style={styles.button}>
  <Text style={styles.text}>Click me</Text>
</Pressable>
```

### After (Gluestack-UI + Tailwind)
```tsx
<Button action="primary" variant="solid" size="md">
  <ButtonText>Click me</ButtonText>
</Button>
```

This provides the same functionality with less code, better consistency, and easier maintenance.