/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.1.dev1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv1

/* 200 ok object */
type GetCorporationsCorporationIdStructures200Ok struct {

	/* ID of the corporation that owns the structure */
	CorporationId int32 `json:"corporation_id,omitempty"`

	/* This week's vulnerability windows, Monday is day 0 */
	CurrentVul []GetCorporationsCorporationIdStructuresCurrentVul `json:"current_vul,omitempty"`

	/* Date on which the structure will run out of fuel */
	FuelExpires SwaggerDateType `json:"fuel_expires,omitempty"`

	/* Next week's vulnerability windows, Monday is day 0 */
	NextVul []GetCorporationsCorporationIdStructuresNextVul `json:"next_vul,omitempty"`

	/* The id of the ACL profile for this citadel */
	ProfileId int32 `json:"profile_id,omitempty"`

	/* Contains a list of service upgrades, and their state */
	Services []GetCorporationsCorporationIdStructuresService `json:"services,omitempty"`

	/* Date at which the structure will move to it's next state */
	StateTimerEnd SwaggerDateType `json:"state_timer_end,omitempty"`

	/* Date at which the structure entered it's current state */
	StateTimerStart SwaggerDateType `json:"state_timer_start,omitempty"`

	/* The Item ID of the structure */
	StructureId int64 `json:"structure_id,omitempty"`

	/* The solar system the structure is in */
	SystemId int32 `json:"system_id,omitempty"`

	/* The type id of the structure */
	TypeId int32 `json:"type_id,omitempty"`

	/* Date at which the structure will unanchor */
	UnanchorsAt SwaggerDateType `json:"unanchors_at,omitempty"`
}
