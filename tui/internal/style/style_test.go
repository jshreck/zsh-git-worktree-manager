package style

import "testing"

func TestCalculateColumnWidths_NormalTerminal(t *testing.T) {
	left, right := CalculateColumnWidths(100)
	if left < MinLeftWidth {
		t.Errorf("left = %d, want >= %d", left, MinLeftWidth)
	}
	if right <= 0 {
		t.Errorf("right = %d, want > 0", right)
	}
	// left + right + 1 (border) should equal total width.
	if left+right+1 != 100 {
		t.Errorf("left(%d) + right(%d) + 1 = %d, want 100", left, right, left+right+1)
	}
}

func TestCalculateColumnWidths_MinTermWidth(t *testing.T) {
	left, right := CalculateColumnWidths(MinTermWidth)
	if left <= 0 {
		t.Errorf("at MinTermWidth, left = %d, want > 0", left)
	}
	if right <= 0 {
		t.Errorf("at MinTermWidth, right = %d, want > 0", right)
	}
}

func TestCalculateColumnWidths_BelowMin(t *testing.T) {
	left, right := CalculateColumnWidths(60)
	if right != 0 {
		t.Errorf("below min width, right = %d, want 0 (single-column mode)", right)
	}
	if left != 60 {
		t.Errorf("below min width, left = %d, want 60", left)
	}
}

func TestCalculateColumnWidths_VeryNarrow(t *testing.T) {
	left, right := CalculateColumnWidths(10)
	if right != 0 {
		t.Errorf("very narrow, right = %d, want 0", right)
	}
	if left != 10 {
		t.Errorf("very narrow, left = %d, want 10", left)
	}
}

func TestCalculateColumnWidths_Wide(t *testing.T) {
	left, right := CalculateColumnWidths(200)
	if left+right+1 != 200 {
		t.Errorf("wide: left(%d) + right(%d) + 1 = %d, want 200", left, right, left+right+1)
	}
}
