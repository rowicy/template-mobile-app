import createClient from "openapi-fetch";
import type { paths } from "@/schema/example";
import { EXAMPLE_API_BASE_URL } from "@/env";

export const client = createClient<paths>({ baseUrl: EXAMPLE_API_BASE_URL });
