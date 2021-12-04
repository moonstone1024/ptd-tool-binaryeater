#!/usr/bin/env bash
set -eu
: "$MD_SRC_DIR"
: "$MD_DEST_DIR"
SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"

decode_md() {
  mkdir -p "$(dirname "$MD_DEST_DIR/$3")"
  "$SCRIPT_DIR/ptd-tool" decode-md "$MD_SRC_DIR/$1" --config "$SCRIPT_DIR/assets/md_loader_configs/$2" | jq > "$MD_DEST_DIR/$3"
  return "$?"
}

mkdir -p ./assets/md/Quests/
decode_md MD_Avatar.snd MD_Avatar_config.json MD_Avatar.json
decode_md MD_Achievement.snd MD_Achievement_config.json MD_Achievement.json
decode_md MD_CharacterMemoryLock.snd MD_CharacterMemoryLock_config.json MD_CharacterMemoryLock.json
decode_md MD_Costume.snd MD_Costume_config.json MD_Costume.json
decode_md MD_SubScenario.snd MD_SubScenario_config.json MD_SubScenario.json
decode_md MD_TowerQuest.snd MD_TowerQuest_config.json MD_TowerQuest.json
decode_md MD_Weapon.snd MD_Weapon_config.json MD_Weapon.json
decode_md MD_Weapon.snd MD_Weapon_config.json MD_Weapon.json
decode_md MD_ScenarioQuest.snd MD_ScenarioQuest_config.json Quests/ScenarioQuest.json
decode_md MD_ScenarioQuest2.snd MD_ScenarioQuest_config.json Quests/ScenarioQuest2.json
decode_md MD_ScenarioQuest3.snd MD_ScenarioQuest_config.json Quests/ScenarioQuest3.json
decode_md MD_ScenarioQuest4.snd MD_ScenarioQuest_config.json Quests/ScenarioQuest4.json
decode_md MD_ScenarioQuest5.snd MD_ScenarioQuest_config.json Quests/ScenarioQuest5.json
decode_md MD_ScenarioQuest6.snd MD_ScenarioQuest_config.json Quests/ScenarioQuest6.json
decode_md MD_GLBQuestc001.snd MD_GLBQuestc_config.json Quests/GLBQuestc001.json
decode_md MD_GLBQuestc002.snd MD_GLBQuestc_config.json Quests/GLBQuestc002.json
decode_md MD_GLBQuestc003.snd MD_GLBQuestc_config.json Quests/GLBQuestc003.json
decode_md MD_GLBQuestc004.snd MD_GLBQuestc_config.json Quests/GLBQuestc004.json
decode_md MD_GLBQuestc005.snd MD_GLBQuestc_config.json Quests/GLBQuestc005.json
decode_md MD_GLBQuestc006.snd MD_GLBQuestc_config.json Quests/GLBQuestc006.json
decode_md MD_GLBQuestc007.snd MD_GLBQuestc_config.json Quests/GLBQuestc007.json
decode_md MD_GLBQuestc008.snd MD_GLBQuestc_config.json Quests/GLBQuestc008.json
decode_md MD_GLBQuestc009.snd MD_GLBQuestc_config.json Quests/GLBQuestc009.json
decode_md MD_ShootingModeItemBackGround.snd MD_ShootingModeItemBackGround_config.json MD_ShootingModeItemBackGround.json
decode_md MD_ShootingModeItemEffect.snd MD_ShootingModeItemEffect_config.json MD_ShootingModeItemEffect.json
decode_md MD_ShootingModeItemFacialExpression.snd MD_ShootingModeItemFacialExpression_config.json MD_ShootingModeItemFacialExpression.json
decode_md MD_ShootingModeItemMotion.snd MD_ShootingModeItemMotion_config.json MD_ShootingModeItemMotion.json
decode_md MD_ShootingModeItemPosing.snd MD_ShootingModeItemPosing_config.json MD_ShootingModeItemPosing.json
decode_md MD_CostumeStone.snd MD_CostumeStone_config.json MD_CostumeStone.json
decode_md MD_WeaponStone.snd MD_WeaponStone_config.json MD_WeaponStone.json
decode_md MD_WeaponSkillContent.snd MD_WeaponSkillContent_config.json MD_WeaponSkillContent.json
